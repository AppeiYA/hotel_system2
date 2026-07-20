INSERT INTO ledger_account (name, type, currency) VALUES
    ('Cash', 'ASSET', 'NGN'),
    ('Room Revenue', 'REVENUE', 'NGN')
ON CONFLICT (name) DO NOTHING;