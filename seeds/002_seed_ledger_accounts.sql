INSERT INTO ledger_account (name, type, currency) VALUES
    ('Guest Receivables', 'ASSET', 'NGN'),
    ('Room Revenue', 'REVENUE', 'NGN'),
    ('Cash', 'ASSET', 'NGN')
ON CONFLICT (name) DO NOTHING;