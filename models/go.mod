module contact_data/models

go 1.17

replace contact_data/persistance => ../persistance

require contact_data/persistance v0.0.0-00010101000000-000000000000

require github.com/lib/pq v1.10.4 // indirect
