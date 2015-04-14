-- Part A
update usgov_hhs_cclf_partas
set claim_payment_amount = claim_payment_amount * -1
where claim_adjustment_type_code = '1';

-- Part A Revenue
update usgov_hhs_cclf_parta_revenues
set claim_line_covered_paid_amount = claim_line_covered_paid_amount * -1
from usgov_hhs_cclf_partas
where usgov_hhs_cclf_partas.current_claim_unique_identifier = usgov_hhs_cclf_parta_revenues.current_claim_unique_identifier
AND usgov_hhs_cclf_partas.claim_adjustment_type_code = '1';

-- Part B
update usgov_hhs_cclf_partb_phys
set claim_line_covered_paid_amount = claim_line_covered_paid_amount * -1
where claim_adjustment_type_code = '1';

--- Part B DME
update usgov_hhs_cclf_partb_dmes
set claim_line_covered_paid_amount = claim_line_covered_paid_amount * -1
where claim_adjustment_type_code = '1';