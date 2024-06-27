-- View: public.submission_user_view

-- DROP VIEW public.submission_user_view;

CREATE OR REPLACE VIEW public.submission_user_view
 AS
 SELECT s.id AS submission_id,
    s.user_id,
    u.name AS user_name,
    u.email AS user_email,
    s.answers,
    s.risk_score,
    s.risk_category,
    s.created_at AS submission_created_at,
    s.updated_at AS submission_updated_at
   FROM submissions s
     JOIN users u ON s.user_id = u.id;

ALTER TABLE public.submission_user_view
    OWNER TO postgres;

