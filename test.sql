CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    discord_id TEXT NOT NULL UNIQUE,
    receive_dm_notifications BOOLEAN DEFAULT 0
);

CREATE TABLE IF NOT EXISTS ethereum_addresses (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    address TEXT,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE IF NOT EXISTS discord_guilds (
    id INTEGER PRIMARY KEY,
    snapshot_space TEXT
);

CREATE TABLE IF NOT EXISTS user_guilds (
    user_id INTEGER,
    guild_id INTEGER,
    PRIMARY KEY (user_id, guild_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (guild_id) REFERENCES discord_guilds (id)
);