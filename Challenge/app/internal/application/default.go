package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jcanonbenavi/app/internal/handler"
	"github.com/jcanonbenavi/app/internal/loader"
	"github.com/jcanonbenavi/app/internal/repository"
	"github.com/jcanonbenavi/app/internal/service"
)

type ConfigAppDefault struct {
	// serverAddr represents the address of the server
	ServerAddr string
	// dbFile represents the path to the database file
	DbFile string
}

type ApplicationDefault struct {
	// router represents the router of the application
	rt *chi.Mux
	// serverAddr represents the address of the server
	serverAddr string
	// dbFile represents the path to the database file
	dbFile string
}

// NewApplicationDefault creates a new default application
func NewApplicationDefault(cfg *ConfigAppDefault) *ApplicationDefault {
	// default values
	defaultRouter := chi.NewRouter()
	defaultConfig := &ConfigAppDefault{
		ServerAddr: ":8080",
		DbFile:     "tickets.csv",
	}
	if cfg != nil {
		if cfg.ServerAddr != "" {
			defaultConfig.ServerAddr = cfg.ServerAddr
		}
		if cfg.DbFile != "" {
			defaultConfig.DbFile = cfg.DbFile
		}
	}
	return &ApplicationDefault{
		rt:         defaultRouter,
		serverAddr: defaultConfig.ServerAddr,
		dbFile:     defaultConfig.DbFile,
	}
}

func (a *ApplicationDefault) SetUp() (err error) {
	// dependencies
	// database
	jsonloader := loader.NewLoaderTicketCSV(a.dbFile)
	db, _ := jsonloader.Load()
	//repository
	rp := repository.NewRepositoryTicketMap(db, 0)
	//service
	service := service.NewServiceTicketDefault(rp)
	//handlers
	handlers := handler.NewDefaulTicket(service)
	//endpoints
	a.rt.Route("/tickets", func(router chi.Router) {
		//get all tickets
		router.Get("/", handlers.Get())
		//get total tickets
		router.Get("/getTotal", handlers.GetTotalTickets())
		//get tickets by country
		router.Get("/getByCountry/{country}", handlers.GetTicketsAmountByDestinationCountry())
		//get average tickets by country
		router.Get("/getAverage/{country}", handlers.GetPercentageTicketsByDestinationCountry())
	})
	err = a.Run()
	return
}

func (a *ApplicationDefault) Run() (err error) {
	// start the server
	//Add a new route to the router
	err = http.ListenAndServe(a.serverAddr, a.rt)
	return
}
