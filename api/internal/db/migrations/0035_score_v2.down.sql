alter table rating_scores
  drop column if exists tier,
  drop column if exists attendance,
  drop column if exists experience,
  drop column if exists gear_score,
  drop column if exists score,
  drop column if exists discipline;
