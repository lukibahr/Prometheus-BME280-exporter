package cmd

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/lukibahr/Prometheus-BME280-exporter/pkg/collectors"
	"github.com/lukibahr/Prometheus-BME280-exporter/pkg/config"
	"github.com/lukibahr/Prometheus-BME280-exporter/pkg/handlers"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shiena/ansicolor"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "prometheus-bme280-exporter",
	Short:        "Export metrics from a Bosh BME280 sensor",
	SilenceUsage: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		return runRoot(cmd)
	},
}

//Execute runs the toor command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("loglevel", "l", "info", "Sets loglevel")
	rootCmd.Flags().StringP("port", "p", "9123", "Sets the port the exporter listens to")
	rootCmd.Flags().StringP("accuracy", "a", "ACCURACY_STANDARD", "Sets the sampling rate of the metric")
}

func setLoglevel(level string) {
	log.SetFormatter(&log.TextFormatter{ForceColors: true})
	log.SetOutput(ansicolor.NewAnsiColorWriter(os.Stdout))
	log.SetReportCaller(false)
	log.SetFormatter(&nested.Formatter{
		HideKeys: false,
	})
	parsed, err := log.ParseLevel(level)
	if err != nil {
		log.Errorf("Invalid loglevel %s", level)
	} else {
		log.SetLevel(parsed)
	}

}

func runRoot(cmd *cobra.Command) error {
	c := collectors.NewBMECollector()
	prometheus.MustRegister(c)

	port, err := cmd.Flags().GetString("port")
	if err != nil {
		return err
	}
	loglevel, err := cmd.Flags().GetString("loglevel")
	if err != nil {
		return err
	}
	accuracy, err := cmd.Flags().GetString("accuracy")
	if err != nil {
		return err
	}
	conf := config.New(port, accuracy, loglevel)
	setLoglevel(conf.Loglevel)
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/health", handlers.HealthHandler)
	http.Handle("/metrics", promhttp.Handler())

	log.Infof("Running exporter on port :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
	return nil

}


