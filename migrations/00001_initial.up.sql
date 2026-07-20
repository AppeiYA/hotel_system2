CREATE TABLE IF NOT EXISTS guest (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    phone TEXT,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS room (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    room_number TEXT NOT NULL UNIQUE,
    room_type TEXT NOT NULL,
    rate BIGINT NOT NULL,
    status TEXT NOT NULL DEFAULT 'available' CHECK (
        status IN ('available','occupied','maintenance')
    ),
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS reservation (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    guest_id UUID NOT NULL REFERENCES guest(id),
    room_id UUID NOT NULL REFERENCES room(id),
    check_in DATE NOT NULL,
    check_out DATE NOT NULL,
    total_amount BIGINT NOT NULL,
    status TEXT NOT NULL DEFAULT 'pending' CHECK (
        status IN ('pending','confirmed','checked_in','checked_out','cancelled', 'no-show')
    ),
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS payment (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    reservation_id UUID NOT NULL REFERENCES reservation(id),
    reference TEXT NOT NULL UNIQUE,
    amount BIGINT NOT NULL,
    method TEXT NOT NULL CHECK (
        method IN ('cash','card','transfer')
    ),
    status TEXT NOT NULL DEFAULT 'pending' CHECK (
        status IN ('pending','success','failed','refunded')
    ),
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS ledger_account (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL UNIQUE,
    type TEXT NOT NULL CHECK (
        type IN ('ASSET','LIABILITY','REVENUE','EXPENSE')
    ),
    currency TEXT NOT NULL DEFAULT 'NGN',
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS ledger_transaction (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    reference_type TEXT NOT NULL,
    reference_id UUID NOT NULL,
    description TEXT,
    status TEXT NOT NULL DEFAULT 'posted' CHECK (
        status IN ('pending','posted','reversed')
    ),
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE IF NOT EXISTS ledger_entry (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    transaction_id UUID NOT NULL REFERENCES ledger_transaction(id),
    account_id UUID NOT NULL REFERENCES ledger_account(id),
    entry_type TEXT NOT NULL CHECK (entry_type IN ('debit','credit')),
    amount BIGINT NOT NULL CHECK (amount > 0),
    created_at TIMESTAMP DEFAULT now()
);

CREATE INDEX idx_reservation_guest ON reservation(guest_id);
CREATE INDEX idx_reservation_room ON reservation(room_id);
CREATE INDEX idx_payment_reservation ON payment(reservation_id);
CREATE INDEX idx_ledger_txn_reference ON ledger_transaction(reference_id);
CREATE INDEX idx_ledger_entries_txn ON ledger_entry(transaction_id);
CREATE INDEX idx_ledger_entries_account ON ledger_entry(account_id);