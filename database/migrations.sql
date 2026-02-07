DROP TABLE IF EXISTS subsystem_state;
-- DROP TABLE IF EXISTS listener_subsystem;
-- DROP TABLE IF EXISTS processor_subsystem;

DROP TABLE IF EXISTS blockchains;
DROP TABLE IF EXISTS blocks;
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS events;

CREATE TABLE IF NOT EXISTS subsystem_state (
    id                     int PRIMARY KEY,
    latest_header_received bigint,
    latest_block_processed bigint
);

CREATE TABLE IF NOT EXISTS blockchains (
    chain_id     int PRIMARY KEY,
    chain_name   varchar(64),
    rpc_endpoint varchar(64),
    latest_block bigint
);

CREATE TABLE IF NOT EXISTS blocks (
    block_number bigint PRIMARY KEY,
    block_hash   varchar(66)
);

CREATE TABLE IF NOT EXISTS transactions (
    trx_hash     varchar(66) PRIMARY KEY,
    block_number bigint
);

CREATE TABLE IF NOT EXISTS events (
    event_id        bigserial PRIMARY KEY,
    event_signature varchar(80),
    trx_hash        varchar(66)
);

-- DROP TABLE IF EXISTS contract_produced_events;
-- DROP TABLE IF EXISTS dao_factory_created_events;
-- DROP TABLE IF EXISTS proposal_factory_created_events;
-- DROP TABLE IF EXISTS erc20_transfer_events;
-- DROP TABLE IF EXISTS erc721_transfer_events;
-- DROP TABLE IF EXISTS member_registered_events;
-- DROP TABLE IF EXISTS proposal_created_events;
-- DROP TABLE IF EXISTS status_updated_events;
-- DROP TABLE IF EXISTS vote_cast_events;
-- DROP TABLE IF EXISTS bay_created_events;
-- DROP TABLE IF EXISTS ask_created_events;
-- DROP TABLE IF EXISTS bid_created_events;
-- DROP TABLE IF EXISTS ask_cancelled_events;
-- DROP TABLE IF EXISTS bid_cancelled_events;
-- DROP TABLE IF EXISTS order_fulfilled_events;

-- CREATE TABLE IF NOT EXISTS contract_produced_events (
--     row_id           serial PRIMARY KEY,
--     event_address    varchar(64),
--     contract_address varchar(64),
--     factory_address  varchar(64),
--     producer         varchar(64)
-- );

-- CREATE TABLE IF NOT EXISTS dao_factory_created_events (
--     row_id           serial PRIMARY KEY,
--     event_address    varchar(64),
--     creator          varchar(64)
-- );

-- CREATE TABLE IF NOT EXISTS proposal_factory_created_events (
--     row_id           serial PRIMARY KEY,
--     event_address    varchar(64),
--     creator          varchar(64)
-- );

-- CREATE TABLE IF NOT EXISTS erc20_transfer_events (
--     row_id           serial PRIMARY KEY,
--     event_address    varchar(64),
--     from_address     varchar(64),
--     to_address       varchar(64),
--     amount           bigint,
--     trx_hash         varchar(64)
-- );

-- CREATE TABLE IF NOT EXISTS erc721_transfer_events (
--     row_id           serial PRIMARY KEY,
--     event_address    varchar(64),
--     from_address     varchar(64),
--     to_address       varchar(64),
--     token_id         bigint,
--     trx_hash         varchar(64)
-- );

-- CREATE TABLE IF NOT EXISTS member_registered_events (
--     row_id           serial PRIMARY KEY,
--     event_address    varchar(64),
--     dao_address      varchar(64),
--     nft_address      varchar(64),
--     token_id         bigint,
--     member_address   varchar(64)
-- );

-- CREATE TABLE IF NOT EXISTS proposal_created_events (
--     row_id           serial PRIMARY KEY,
--     event_address    varchar(64),
--     dao_address      varchar(64),
--     proposal_address varchar(64),
--     proposal_id      bigint,
--     creator          varchar(64)
-- );

-- CREATE TABLE IF NOT EXISTS status_updated_events (
--     row_id           serial PRIMARY KEY,
--     event_address    varchar(64),
--     new_status       varchar(32)
-- );

-- CREATE TABLE IF NOT EXISTS vote_cast_events (
--     row_id           serial PRIMARY KEY,
--     event_address    varchar(64),
--     voter            varchar(64),
--     token_id         bigint,
--     vote             int
-- );

-- CREATE TABLE IF NOT EXISTS bay_created_events (
--     row_id           serial PRIMARY KEY,
--     event_address    varchar(64),
--     dao_address      varchar(64),
--     admin_address    varchar(64)
-- );

-- CREATE TABLE IF NOT EXISTS ask_created_events (
--     row_id           serial PRIMARY KEY,
--     event_address    varchar(64),
--     order_hash       varchar(128),
--     creator          varchar(64),
--     erc721_address   varchar(64),
--     token_id         bigint,
--     erc20_address    varchar(64),
--     amount           bigint
-- );

-- CREATE TABLE IF NOT EXISTS bid_created_events (
--     row_id           serial PRIMARY KEY,
--     event_address    varchar(64),
--     order_hash       varchar(128),
--     creator          varchar(64),
--     erc721_address   varchar(64),
--     token_id         bigint,
--     erc20_address    varchar(64),
--     amount           bigint
-- );

-- CREATE TABLE IF NOT EXISTS ask_cancelled_events (
--     row_id           serial PRIMARY KEY,
--     event_address    varchar(64),
--     order_hash       varchar(128)
-- );

-- CREATE TABLE IF NOT EXISTS bid_cancelled_events (
--     row_id           serial PRIMARY KEY,
--     event_address    varchar(64),
--     order_hash       varchar(128)
-- );

-- CREATE TABLE IF NOT EXISTS order_fulfilled_events (
--     row_id           serial PRIMARY KEY,
--     event_address    varchar(64),
--     order_hash       varchar(128)
-- );
