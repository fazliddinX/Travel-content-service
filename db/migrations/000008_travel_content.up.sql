CREATE TABLE destinations (
                              id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                              name VARCHAR(100) NOT NULL,
                              country VARCHAR(100) NOT NULL,
                              description TEXT,
                              best_time_to_visit VARCHAR(100),
                              average_cost_per_day DECIMAL(10, 2),
                              currency VARCHAR(3),
                              language VARCHAR(50),
                              created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                              updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                              deleted_at  TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
