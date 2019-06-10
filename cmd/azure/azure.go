package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/ghodss/yaml"
	"github.com/kelseyhightower/envconfig"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig"
	"github.com/openshift/installer/pkg/asset/store"

	targetassets "github.com/openshift/installer/pkg/asset/targets"
	"github.com/openshift/installer/pkg/ipnet"
	"github.com/openshift/installer/pkg/types"
	"github.com/openshift/installer/pkg/types/azure"
	"github.com/openshift/installer/pkg/types/defaults"
	"github.com/openshift/installer/pkg/types/validation"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/openshift/openshift-azure/pkg/util/random"
)

var (
	gitCommit  = "unknown"
	baseDomain = "osadev.cloud"
)

type EnvConfig struct {
	Region  string
	Regions string `envconfig:"AZURE_REGIONS" required:"true"`
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func newEnvConfig() (*EnvConfig, error) {
	var c EnvConfig
	if err := envconfig.Process("", &c); err != nil {
		return nil, err
	}
	regions := strings.Split(c.Regions, ",")
	rand.Seed(time.Now().UTC().UnixNano())
	c.Region = regions[rand.Intn(len(regions))]
	if c.Region == "" {
		return nil, fmt.Errorf("must set AZURE_REGIONS to a comma separated list")
	}
	return &c, nil
}

func getInstallConfig(name string) (*types.InstallConfig, error) {
	// TODO: move to util/secrets
	fqdn, err := random.FQDN(baseDomain, 5)
	if err != nil {
		return nil, err
	}

	file, err := os.Open("secrets/pull-secret.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	pullSecret, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	ec, err := newEnvConfig()
	if err != nil {
		return nil, err
	}

	cfg := types.InstallConfig{
		TypeMeta: metav1.TypeMeta{
			APIVersion: types.InstallConfigVersion,
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		BaseDomain: fqdn,
		Compute: []types.MachinePool{
			{
				Name:           "worker",
				Replicas:       to.Int64Ptr(1),
				Hyperthreading: types.HyperthreadingEnabled,
			},
		},
		Networking: &types.Networking{
			MachineCIDR:    ipnet.MustParseCIDR("10.0.0.0/16"),
			NetworkType:    "OpenShiftSDN",
			ServiceNetwork: []ipnet.IPNet{*ipnet.MustParseCIDR("172.30.0.0/16")},
			ClusterNetwork: []types.ClusterNetworkEntry{
				{
					CIDR:       *ipnet.MustParseCIDR("10.128.0.0/14"),
					HostPrefix: 23,
				},
			},
		},
		ControlPlane: &types.MachinePool{
			Name:           "master",
			Replicas:       to.Int64Ptr(3),
			Hyperthreading: types.HyperthreadingEnabled,
		},
		Platform: types.Platform{
			Azure: &azure.Platform{
				Region: ec.Region,
			},
		},
		PullSecret: string(pullSecret),
	}
	return &cfg, nil
}

func run() error {
	rootCmd := &cobra.Command{
		Use:  "./azure [component]",
		Long: "Azure Red Hat OpenShift dispatcher",
	}
	rootCmd.PersistentFlags().StringP("loglevel", "l", "Debug", "Valid values are [Debug, Info, Warning, Error]")
	rootCmd.PersistentFlags().StringP("name", "n", "demo-cluster", "Cluster name value")
	rootCmd.Printf("gitCommit %s\n", gitCommit)
	rootCmd.Execute()

	name, err := rootCmd.Flags().GetString("name")
	if err != nil {
		return err
	}

	cfg, err := getInstallConfig(name)
	if err != nil {
		return errors.Wrap(err, "failed to get InstallConfig")
	}

	defaults.SetInstallConfigDefaults(cfg)
	if err := validation.ValidateInstallConfig(cfg, openstackvalidation.NewValidValuesFetcher()).ToAggregate(); err != nil {
		return errors.Wrap(err, "invalid install config")
	}
	fmt.Println(cfg)

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return errors.Wrap(err, "failed to Marshal InstallConfig")
	}

	// doing this to prevent the stdin questions.
	ic := &installconfig.InstallConfig{}
	ic.Config = cfg
	ic.File = &asset.File{
		Filename: "install-config.yaml",
		Data:     data,
	}
	waIc := []asset.WritableAsset{ic}

	// TODO: Start consuming assets. Assets need to be file system agnostic
	// Like: https://github.com/openshift/installer/blob/master/pkg/asset/installconfig/ssh.go#L49

	// TODO re-implement the store to not use the filesystem?
	directory := path.Join("clusters", name)
	assetStore, err := store.NewStore(directory)
	if err != nil {
		return errors.Wrap(err, "failed to create asset store")
	}

	targets := waIc
	targets = append(targets, targetassets.IgnitionConfigs...)
	targets = append(targets, targetassets.Manifests...)
	targets = append(targets, targetassets.Cluster...)

	for _, a := range targets {
		err := assetStore.Fetch(a, targets...)
		if err != nil {
			err = errors.Wrapf(err, "failed to fetch %s", a.Name())
		}

		if err2 := asset.PersistToFile(a, directory); err2 != nil {
			err2 = errors.Wrapf(err2, "failed to write asset (%s) to disk", a.Name())
			if err != nil {
				return err
			}
			return err2
		}

		if err != nil {
			return err
		}
	}

	// wait for the cluster to come up
	//	err = waitForBootstrapComplete(ctx, config, rootOpts.dir)
	//	err = destroybootstrap.Destroy(rootOpts.dir)
	//	err = waitForInstallComplete(ctx, config, rootOpts.dir)

	return nil
}
