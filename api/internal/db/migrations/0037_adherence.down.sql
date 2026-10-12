-- Back to 0036: the rows keep their last-raid facts, without adherence.
alter table rating_scores drop column if exists adherence;
