CREATE TABLE itinerary_destinations (
                                        id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                        itinerary_id UUID REFERENCES itineraries(id),
                                        name VARCHAR(100) NOT NULL,
                                        start_date DATE NOT NULL,
                                        end_date DATE NOT NULL
);
