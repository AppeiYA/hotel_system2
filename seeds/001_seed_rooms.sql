INSERT INTO room (room_number, room_type, rate, status) VALUES
    ('101', 'single',      1500000, 'available'),
    ('102', 'single',      1500000, 'available'),
    ('103', 'single',      1500000, 'occupied'),
    ('104', 'single',      1500000, 'maintenance'),
    ('201', 'double',      2500000, 'available'),
    ('202', 'double',      2500000, 'available'),
    ('203', 'double',      2500000, 'occupied'),
    ('204', 'double',      2500000, 'available'),
    ('301', 'deluxe',      4000000, 'available'),
    ('302', 'deluxe',      4000000, 'occupied'),
    ('303', 'deluxe',      4000000, 'available'),
    ('401', 'suite',       7500000, 'available'),
    ('402', 'suite',       7500000, 'maintenance'),
    ('501', 'presidential', 15000000, 'available')
ON CONFLICT (room_number) DO NOTHING;