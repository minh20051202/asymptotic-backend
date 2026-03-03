package catalog

import (
	"database/sql"
)

type CatalogRepository interface {
	Init() error

	CreateProvider(provider *Provider) error
	GetProviderByName(providerName string) (*Provider, error)

	CreateEndpoint(endpoint *Endpoint) error
	GetEndpointById(endpointId int64) (*Endpoint, error)

	CreatePricing(pricing *Pricing) error
	GetActivePricingByEndpointId(endpointId int64) (*Pricing, error)
	SetNewPrice(endpointId int64, newCost int64) error

	GetRouteConfigById(endpointId int64) (*RouteConfig, error)
	GetRouteConfigByNames(providerName, endpointName, endpointPath string) (*RouteConfig, error)
	GetAllAvailableTools() ([]*ToolInfo, error)
}

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Init() error {
	if err := r.createProviderTable(); err != nil {
		return err
	}
	if err := r.createEndpointTable(); err != nil {
		return err
	}
	if err := r.createPricingTable(); err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) createProviderTable() error {
	query := `CREATE TABLE IF NOT EXISTS providers (
    	provider_id SERIAL PRIMARY KEY,
   	 	name VARCHAR(50) UNIQUE NOT NULL,
    	base_url VARCHAR(255) UNIQUE NOT NULL,
    	encrypted_api_key TEXT NOT NULL,
    	is_active BOOLEAN DEFAULT TRUE,
    	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) createEndpointTable() error {
	query := `CREATE TABLE IF NOT EXISTS endpoints (
		endpoint_id SERIAL PRIMARY KEY,
		provider_id INT NOT NULL,
		name VARCHAR(100) NOT NULL,
		http_method VARCHAR(10) NOT NULL,
		path VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		
		UNIQUE(provider_id, http_method, path),
		
		CONSTRAINT fk_endpoint_provider
			FOREIGN KEY (provider_id) 
			REFERENCES providers(provider_id)
			ON DELETE RESTRICT
	)`
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) createPricingTable() error {
	query := `CREATE TABLE IF NOT EXISTS pricings (
		pricing_id SERIAL PRIMARY KEY,
		endpoint_id INT NOT NULL,
		cost BIGINT NOT NULL,
		is_current BOOLEAN DEFAULT TRUE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		valid_until TIMESTAMP,

		CONSTRAINT fk_pricing_endpoint
			FOREIGN KEY (endpoint_id) 
			REFERENCES endpoints(endpoint_id)
			ON DELETE RESTRICT
	);`
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository) CreateProvider(provider *Provider) error {
	query := `
		INSERT INTO providers (name, base_url, encrypted_api_key, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5) 
	`
	_, err := r.db.Exec(
		query,
		provider.Name,
		provider.BaseUrl,
		provider.EncryptedApiKey,
		provider.CreatedAt,
		provider.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) GetProviderByName(providerName string) (*Provider, error) {
	provider := new(Provider)

	query := `
		SELECT provider_id, name, base_url, encrypted_api_key, is_active, created_at, updated_at 
		FROM providers 
		WHERE name = $1
	`

	err := r.db.QueryRow(query, providerName).Scan(
		&provider.ProviderId,
		&provider.Name,
		&provider.BaseUrl,
		&provider.EncryptedApiKey,
		&provider.IsActive,
		&provider.CreatedAt,
		&provider.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return provider, nil
}

func (r *PostgresRepository) CreateEndpoint(endpoint *Endpoint) error {
	queryInsert := `
		INSERT INTO endpoints (provider_id, name, http_method, path, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(
		queryInsert,
		endpoint.ProviderId,
		endpoint.Name,
		endpoint.HttpMethod,
		endpoint.Path,
		endpoint.CreatedAt,
		endpoint.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
func (r *PostgresRepository) GetEndpointById(endpointId int64) (*Endpoint, error) {
	endpoint := new(Endpoint)

	query := `
		SELECT endpoint_id, provider_id, name, http_method, path, created_at, updated_at 
		FROM endpoints
		WHERE endpoint_id = $1
	`

	err := r.db.QueryRow(query, endpointId).Scan(
		&endpoint.EndpointId,
		&endpoint.ProviderId,
		&endpoint.Name,
		&endpoint.HttpMethod,
		&endpoint.Path,
		&endpoint.CreatedAt,
		&endpoint.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return endpoint, nil
}

func (r *PostgresRepository) CreatePricing(pricing *Pricing) error {
	queryInsert := `
		INSERT INTO pricings (endpoint_id, cost, created_at) 
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(
		queryInsert,
		pricing.EndpointId,
		pricing.Cost,
		pricing.CreatedAt,
	)

	if err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) GetActivePricingByEndpointId(endpointId int64) (*Pricing, error) {
	pricing := new(Pricing)

	query := `
		SELECT pricing_id, endpoint_id, cost, is_current, created_at, valid_until
		FROM pricings
		WHERE endpoint_id = $1 AND is_current = TRUE
	`

	err := r.db.QueryRow(query, endpointId).Scan(
		&pricing.PricingId,
		&pricing.EndpointId,
		&pricing.Cost,
		&pricing.IsCurrent,
		&pricing.CreatedAt,
		&pricing.ValidUntil,
	)

	if err != nil {
		return nil, err
	}

	return pricing, nil
}

func (r *PostgresRepository) SetNewPrice(endpointId int64, newCost int64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var currentPricingId int64
	queryRead := `
		SELECT pricing_id 
		FROM pricings 
		WHERE endpoint_id = $1 AND is_current = TRUE 
		FOR UPDATE
	`
	err = tx.QueryRow(queryRead, endpointId).Scan(&currentPricingId)

	if err != nil {
		return err
	}

	queryUpdate := `
		UPDATE pricings 
		SET is_current = FALSE, valid_until = CURRENT_TIMESTAMP 
		WHERE pricing_id = $1
	`
	_, err = tx.Exec(queryUpdate, currentPricingId)

	if err != nil {
		return err
	}

	queryInsert := `
		INSERT INTO pricings (endpoint_id, cost) 
		VALUES ($1, $2)
	`

	_, err = tx.Exec(queryInsert, endpointId, newCost)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *PostgresRepository) GetRouteConfigById(endpointId int64) (*RouteConfig, error) {
	config := new(RouteConfig)

	query := `
		SELECT P1.base_url, P1.encrypted_api_key, E.path, E.http_method, P2.cost 
		FROM providers P1
		INNER JOIN endpoints E ON P1.provider_id = E.provider_id
		INNER JOIN pricings P2 ON P2.endpoint_id = E.endpoint_id
		WHERE P1.is_active = TRUE 
		  AND P2.is_current = TRUE
		  AND E.endpoint_id = $1
	`

	err := r.db.QueryRow(query, endpointId).Scan(
		&config.ProviderBaseUrl,
		&config.EncryptedApiKey,
		&config.EndpointPath,
		&config.EndpointHttpMethod,
		&config.Cost,
	)

	if err != nil {
		return nil, err
	}

	return config, nil
}

func (r *PostgresRepository) GetRouteConfigByNames(providerName, endpointName, endpointPath string) (*RouteConfig, error) {
	config := new(RouteConfig)

	query := `
		SELECT P1.base_url, P1.encrypted_api_key, E.path, E.http_method, P2.cost 
		FROM providers P1
		INNER JOIN endpoints E ON P1.provider_id = E.provider_id
		INNER JOIN pricings P2 ON P2.endpoint_id = E.endpoint_id
		WHERE P1.is_active = TRUE 
		  AND P2.is_current = TRUE
		  AND P1.name = $1
		  AND E.name = $2
		  AND E.path = $3
	`

	err := r.db.QueryRow(query, providerName, endpointName, endpointPath).Scan(
		&config.ProviderBaseUrl,
		&config.EncryptedApiKey,
		&config.EndpointPath,
		&config.EndpointHttpMethod,
		&config.Cost,
	)

	if err != nil {
		return nil, err
	}

	return config, nil
}

func (r *PostgresRepository) GetAllAvailableTools() ([]*ToolInfo, error) {
	query := `
		SELECT E.endpoint_id, P1.name, E.name, E.http_method, E.path, P2.cost 
		FROM providers P1
		INNER JOIN endpoints E ON P1.provider_id = E.provider_id
		INNER JOIN pricings P2 ON P2.endpoint_id = E.endpoint_id
		WHERE P1.is_active = TRUE 
		  AND P2.is_current = TRUE
	`

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	tools := []*ToolInfo{}

	for rows.Next() {
		tool, err := scanIntoToolInfos(rows)
		if err != nil {
			return nil, err
		}
		tools = append(tools, tool)
	}
	return tools, nil
}

func scanIntoToolInfos(rows *sql.Rows) (*ToolInfo, error) {
	tool := new(ToolInfo)
	err := rows.Scan(
		&tool.EndpointId,
		&tool.ProviderName,
		&tool.EndpointName,
		&tool.HttpMethod,
		&tool.Path,
		&tool.Cost,
	)
	return tool, err
}
