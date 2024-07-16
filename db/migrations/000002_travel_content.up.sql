CREATE TABLE story_tags (
                            story_id UUID DEFAULT gen_random_uuid() REFERENCES stories(id),
                            tag VARCHAR(50),
                            PRIMARY KEY (story_id, tag)
);