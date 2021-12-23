package app

import (
	"context"
	"fmt"
	"github.com/gobuffalo/logger"
	"github.com/markbates/oncer"
	"github.com/mrNobody95/Gate/services/dolphin/configs"
	"github.com/mrNobody95/Gate/services/dolphin/internal/pkg/defaults"
	"github.com/sirupsen/logrus"
	"net/http"
)

// Options are used to configure and define how your application should run.
type Options struct {
	Name string `json:"name"`
	// Addr is the bind address provided to http.Server. Default is "127.0.0.1:3000"
	// Can be set using ENV vars "ADDR" and "PORT".
	Addr string `json:"addr"`
	// Host that this application will be available at. Default is "http://127.0.0.1:[$PORT|3000]".
	Host string `json:"host"`

	// Env is the "environment" in which the App is running. Default is "development".
	Env string `json:"env"`

	// LogLevel defaults to "debug". Deprecated use LogLvl instead
	LogLevel string `json:"log_level"`
	// LogLevl defaults to logger.DebugLvl.
	LogLvl logger.Level `json:"log_lvl"`
	// Logger to be used with the application. A default one is provided.
	Logger       *logrus.Logger `json:"-"`
	AccessLogger *logrus.Logger `json:"-"`

	// MethodOverride allows for changing of the request method type. See the default
	// implementation at buffalo.MethodOverride
	MethodOverride http.HandlerFunc `json:"-"`

	//// SessionStore is the `github.com/gorilla/sessions` store used to back
	//// the session. It defaults to use a cookie store and the ENV variable
	//// `SESSION_SECRET`.
	//SessionStore sessions.Store `json:"-"`
	// SessionName is the name of the session cookie that is set. This defaults
	// to "_buffalo_session".
	SessionName string `json:"session_name"`

	//// Worker implements the Worker interface and can process tasks in the background.
	//// Default is "github.com/gobuffalo/worker.Simple.
	//Worker worker.Worker `json:"-"`
	// WorkerOff tells App.Start() whether to start the Worker process or not. Default is "false".
	WorkerOff bool `json:"worker_off"`

	// PreHandlers are http.Handlers that are called between the http.Server
	// and the buffalo Application.
	PreHandlers []http.Handler `json:"-"`
	// PreWare takes an http.Handler and returns and http.Handler
	// and acts as a pseudo-middleware between the http.Server and
	// a Buffalo application.
	PreWares []PreWare `json:"-"`

	Prefix  string          `json:"prefix"`
	Context context.Context `json:"-"`

	cancel context.CancelFunc
}

// PreWare takes an http.Handler and returns and http.Handler
// and acts as a pseudo-middleware between the http.Server and
// a Buffalo application.
type PreWare func(http.Handler) http.Handler

// NewOptions returns a new Options instance with sensible defaults
func NewOptions() Options {
	return optionsWithDefaults(Options{})
}

func optionsWithDefaults(opts Options) Options {
	opts.Env = defaults.String(opts.Env, configs.Variables.Environment)
	opts.Name = defaults.String(opts.Name, "/")
	envAddr := configs.Variables.Host

	opts.Addr = defaults.String(opts.Addr, fmt.Sprintf("%s:%d", envAddr, configs.Variables.Port))

	if opts.PreWares == nil {
		opts.PreWares = []PreWare{}
	}
	if opts.PreHandlers == nil {
		opts.PreHandlers = []http.Handler{}
	}

	if opts.Context == nil {
		opts.Context = context.Background()
	}
	opts.Context, opts.cancel = context.WithCancel(opts.Context)

	if opts.Logger == nil {
		var err error
		opts.LogLvl, err = logger.ParseLevel(configs.Variables.LogLevel)
		if err != nil {
			opts.LogLvl = logger.DebugLevel
		}

		if len(opts.LogLevel) > 0 {
			var err error
			oncer.Deprecate(0, "github.com/gobuffalo/buffalo#Options.LogLevel", "Use github.com/gobuffalo/buffalo#Options.LogLvl instead.")
			opts.LogLvl, err = logger.ParseLevel(opts.LogLevel)
			if err != nil {
				opts.LogLvl = logger.DebugLevel
			}
		}
		if opts.LogLvl == 0 {
			opts.LogLvl = logger.DebugLevel
		}
		opts.Logger = logrus.StandardLogger()
		//opts.Logger = logger.New(opts.LogLvl)
	}
	if opts.Logger == nil {
		opts.AccessLogger = logrus.StandardLogger()
	}

	//pop.SetLogger(func(level logging.Level, s string, args ...interface{}) {
	//	if !pop.Debug {
	//		return
	//	}
	//
	//	l := opts.Logger
	//	if len(args) > 0 {
	//		for i, a := range args {
	//			l = l.WithField(fmt.Sprintf("$%d", i+1), a)
	//		}
	//	}
	//
	//	if pop.Color {
	//		s = color.YellowString(s)
	//	}
	//
	//	l.Debug(s)
	//})

	//if opts.SessionStore == nil {
	//	secret := envy.Get("SESSION_SECRET", "") change with my env
	//
	//	if secret == "" && (opts.Env == "development" || opts.Env == "test") {
	//		secret = "buffalo-secret"
	//	}
	//
	//	// In production a SESSION_SECRET must be set!
	//	if secret == "" {
	//		opts.Logger.Warn("Unless you set SESSION_SECRET env variable, your session storage is not protected!")
	//	}
	//
	//	cookieStore := sessions.NewCookieStore([]byte(secret))
	//
	//	//Cookie secure attributes, see: https://www.owasp.org/index.php/Testing_for_cookies_attributes_(OTG-SESS-002)
	//	cookieStore.Options.HttpOnly = true
	//	if opts.Env == "production" {
	//		cookieStore.Options.Secure = true
	//	}
	//
	//	opts.SessionStore = cookieStore
	//}
	//if opts.Worker == nil {
	//	w := worker.NewSimpleWithContext(opts.Context)
	//	w.Logger = opts.Logger
	//	opts.Worker = w
	//}
	opts.SessionName = defaults.String(opts.SessionName, "_buffalo_session")
	scheme := "http"
	if configs.Variables.TLS {
		scheme = "https"
	}
	opts.Host = defaults.String(opts.Host, fmt.Sprintf("%s://%s:%d", scheme, configs.Variables.Host, configs.Variables.Port))
	return opts
}