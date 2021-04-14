package server

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
)

func Init(log *logrus.Entry) {
	// Set server logger
	setLogger(log)

	loadConfigFromEnvs()
	loadAuthConfigFromEnvs()
	// logger.Info("Server succesfully started")
}

func Start() error {
	router := NewRouter(logger)

	httpServer := http.Server{
		Addr:           fmt.Sprintf("%s:%d", configuration.Host, configuration.Port),
		Handler:        router,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	switch configuration.SSL {
	case true:
		tlsConfig := &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305, // Go 1.8 only
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,   // Go 1.8 only
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,

				// Best disabled, as they don't provide Forward Secrecy,
				// but might be necessary for some clients
				// tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				// tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
			},
			PreferServerCipherSuites: true,
			MinVersion:               tls.VersionTLS12,
			CurvePreferences: []tls.CurveID{
				tls.CurveP256,
				tls.X25519, // Go 1.8 only
			},
			InsecureSkipVerify: true,
		}

		httpServer.TLSConfig = tlsConfig

		logger.Infof("Listening on port %s:%d", configuration.Host, configuration.Port)

		if err := httpServer.ListenAndServeTLS(configuration.ServerCertPath, configuration.ServerKeyPath); err != nil {
			return errors.Wrap(err, "failed to start https server")
		}

	default:
		logger.Infof("Listening on port %s:%d", configuration.Host, configuration.Port)

		if err := httpServer.ListenAndServe(); err != nil {
			return errors.Wrap(err, "failed to start http server")
		}
	}

	return nil
}
