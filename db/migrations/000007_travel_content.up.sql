CREATE TABLE itinerary_activities (
                                      id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                      destination_id UUID REFERENCES itinerary_destinations(id),
                                      activity TEXT NOT NULL
);
