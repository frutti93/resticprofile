package prom

import (
	"github.com/prometheus/client_golang/prometheus"
)

type Status uint

const (
	StatusFailed Status = iota
	StatusWarning
	StatusSuccess
)

type BackupMetrics struct {
	duration         *prometheus.GaugeVec
	filesNew         *prometheus.GaugeVec
	filesChanged     *prometheus.GaugeVec
	filesUnmodified  *prometheus.GaugeVec
	dirNew           *prometheus.GaugeVec
	dirChanged       *prometheus.GaugeVec
	dirUnmodified    *prometheus.GaugeVec
	filesTotal       *prometheus.GaugeVec
	bytesAdded       *prometheus.GaugeVec
	bytesAddedPacked *prometheus.GaugeVec
	bytesTotal       *prometheus.GaugeVec
	status           *prometheus.GaugeVec
	time             *prometheus.GaugeVec
}

func newBackupMetrics(labels []string) BackupMetrics {
	backupMetrics := BackupMetrics{
		duration: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: backup,
			Name:      "duration_seconds",
			Help:      "The backup duration (in seconds).",
		}, labels),
		filesNew: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: backup,
			Name:      "files_new",
			Help:      "Number of new files added to the backup.",
		}, labels),
		filesChanged: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: backup,
			Name:      "files_changed",
			Help:      "Number of files with changes.",
		}, labels),
		filesUnmodified: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: backup,
			Name:      "files_unmodified",
			Help:      "Number of files unmodified since last backup.",
		}, labels),
		dirNew: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: backup,
			Name:      "dir_new",
			Help:      "Number of new directories added to the backup.",
		}, labels),
		dirChanged: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: backup,
			Name:      "dir_changed",
			Help:      "Number of directories with changes.",
		}, labels),
		dirUnmodified: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: backup,
			Name:      "dir_unmodified",
			Help:      "Number of directories unmodified since last backup.",
		}, labels),
		filesTotal: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: backup,
			Name:      "files_processed",
			Help:      "Total number of files scanned by the backup for changes.",
		}, labels),
		bytesAdded: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: backup,
			Name:      "added_bytes",
			Help:      "Total number of bytes added to the repository.",
		}, labels),
		bytesAddedPacked: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: backup,
			Name:      "added_bytes_packed",
			Help:      "Total number of bytes added to the repository after compression.",
		}, labels),
		bytesTotal: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: backup,
			Name:      "processed_bytes",
			Help:      "Total number of bytes scanned for changes.",
		}, labels),
		status: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: backup,
			Name:      "status",
			Help:      "Backup status: 0=fail, 1=warning, 2=success.",
		}, labels),
		time: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: backup,
			Name:      "time_seconds",
			Help:      "Last backup run (unixtime).",
		}, labels),
	}
	return backupMetrics
}
