CREATE TABLE IF NOT EXISTS trips (
    id UUID PRIMARY KEY,
    destination TEXT NOT NULL,
    start_date DATE,
    end_date DATE,
    owner_name TEXT NOT NULL,
    owner_email TEXT NOT NULL,
    status INTEGER -- 1 para verdadeiro (true), 0 para falso (false)
);

CREATE TABLE IF NOT EXISTS emails_to_invite (
    id UUID PRIMARY KEY,
    trip_id UUID,
    email TEXT NOT NULL,
    FOREIGN KEY (trip_id) REFERENCES trips(id)
);

CREATE TABLE IF NOT EXISTS links (
    id UUID PRIMARY KEY,
    trip_id UUID,
    link TEXT NOT NULL,
    title TEXT NOT NULL,
    FOREIGN KEY (trip_id) REFERENCES trips(id)
);

CREATE TABLE IF NOT EXISTS participants (
    id UUID PRIMARY KEY,
    trip_id UUID NOT NULL,
    emails_to_invite_id UUID NOT NULL,
    name TEXT NOT NULL,
    is_confirmed INTEGER, -- 1 para verdadeiro (true), 0 para falso (false)
    FOREIGN KEY (trip_id) REFERENCES trips(id),
    FOREIGN KEY (emails_to_invite_id) REFERENCES emails_to_invite(id)
);

CREATE TABLE IF NOT EXISTS activities (
    id UUID PRIMARY KEY,
    trip_id UUID NOT NULL,
    title TEXT NOT NULL,
    occurs_at TIMESTAMP,
    FOREIGN KEY (trip_id) REFERENCES trips(id)
);