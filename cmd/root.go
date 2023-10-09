package cmd

import (
	"net/http"
	"os"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/lukibahr/Prometheus-BME280-exporter/pkg/collectors"
	"github.com/lukibahr/Prometheus-BME280-exporter/pkg/config"
	"github.com/lukibahr/Prometheus-BME280-exporter/pkg/handlers"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shiena/ansicolor"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	buildVersion string
	buildCommit  string
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

// Execute runs the toor command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("loglevel", "l", "info", "Sets loglevel")
	rootCmd.Flags().StringP("port", "p", "9123", "Sets the port the exporter listens to")
	rootCmd.Flags().StringP("accuracy", "a", "ACCURACY_STANDARD", "Sets the sampling rate of the metric")
	rootCmd.Flags().StringP("mqttbroker", "b", "localhost", "set the mqtt broker hostname")
	rootCmd.Flags().StringP("mqttport", "o", "1883", "set the mqtt broker port")
	rootCmd.Flags().StringP("mqttusername", "u", "dev", "set the mqtt broker username")
	rootCmd.Flags().StringP("mqttpassword", "w", "dev", "set the mqtt broker password")
	rootCmd.Flags().StringP("mode", "m", "prometheus", "enable prometheus exporter or mqtt")
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
		log.Errorf("invalid loglevel %s", level)
	} else {
		log.SetLevel(parsed)
	}
}

func runRoot(cmd *cobra.Command) error {
	log.Infof("prometheus-bme280-exporter version %s, commit %s", buildVersion, buildCommit)

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
	mode, err := cmd.Flags().GetString("mode")
	if err != nil {
		return err
	}
	mqttbroker, err := cmd.Flags().GetString("mqttbroker")
	if err != nil {
		return err
	}
	mqttport, err := cmd.Flags().GetString("mqttport")
	if err != nil {
		return err
	}
	mqttusername, err := cmd.Flags().GetString("mqttusername")
	if err != nil {
		return err
	}
	mqttpassword, err := cmd.Flags().GetString("mqttpassword")
	if err != nil {
		return err
	}

	conf := config.New(port, accuracy, loglevel, mqttbroker, mqttport, mqttusername, mqttpassword)
	setLoglevel(conf.Loglevel)

	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/health", handlers.HealthHandler)
	switch mode {
	case "prometheus":
		log.Info("running with mode prometheus to serve as an exporter")
		c := collectors.NewBMECollector()
		prometheus.MustRegister(c)
		http.Handle("/metrics", promhttp.Handler())
		log.Infof("Running exporter on port :%s", port)
		log.Fatal(http.ListenAndServe(":"+port, nil))
	case "mqtt":
		log.Info("running with mode mqtt")
		collectors.PubSub(mqttbroker, mqttport, mqttusername, mqttpassword)
	}

	return nil

}
