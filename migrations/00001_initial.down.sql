DROP INDEX idx_reservation_guest;
DROP INDEX idx_reservation_room;
DROP INDEX idx_payment_reservation;
DROP INDEX idx_ledger_txn_reference;
DROP INDEX idx_ledger_entries_txn;
DROP INDEX idx_ledger_entries_account;

DROP TABLE IF EXISTS ledger_entry;
DROP TABLE IF EXISTS ledger_transaction;
DROP TABLE IF EXISTS ledger_account;
DROP TABLE IF EXISTS payment;
DROP TABLE IF EXISTS reservation;
DROP TABLE IF EXISTS room;
DROP TABLE IF EXISTS guest;