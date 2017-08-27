do $$
declare
  alpha_uuid uuid   := 'c9c32c6d-5f1c-444f-b185-93c358a2eb36';
  beta_uuid uuid    := 'd0ac0dee-670c-4524-a716-d60485e31976';
  gamma_uuid uuid   := '746278e0-f497-490a-814f-25d685e4db3a';
  delta_uuid uuid   := '88dafdad-d158-4b19-bcf8-7daca5b23481';
  epsilon_uuid uuid := '5a964485-d9f4-4080-afd5-0cd90f80b832';
  zeta_uuid uuid    := '70992e6b-ce49-47ce-b52c-8f69e28c578e';
  phi_uuid uuid     := 'cad4335d-0ca3-4d79-aa10-6a168b46e1d0';
  chi_uuid uuid     := 'f680ebdd-678e-4c74-b05b-045fb0e20e1e';
  psi_uuid uuid     := 'a30035d6-c9e2-4376-8773-954384830509';
  omega_uuid uuid   := 'f2419558-940a-47e1-aefa-075e27a078b3';
  eta_uuid uuid     := 'd325ef73-7eab-4c8b-8333-28bb0187a467';
  theta_uuid uuid   := '572dd0d7-32d5-4234-a7f7-42933c3d4b5a';
  iota_uuid uuid    := '082941d4-f5f3-4e2e-b37d-ea7b4b346d87';
begin

  insert into projects
    (uuid, user_id, repository, state_cd, supported_swift_version)
  values
    (alpha_uuid, null, 'org/alpha-ios', 0, null),
    (beta_uuid, null, 'org/beta-ios', 1, null),
    (gamma_uuid, null, 'org/gamma-ios', 2, null),
    (delta_uuid, null, 'org/delta-ios', 3, '2.0'),
    (epsilon_uuid, null, 'org/epsilon-ios', 4, null),
    (zeta_uuid, 1, 'test1/zeta-ios', 3, '2.0'),
    (eta_uuid, 1, 'test1/eta-ios', 3, '3.0'),
    (theta_uuid, 2, 'test2/theta-ios', 3, '2.0'),
    (iota_uuid, 2, 'test2/iota-ios', 3, '3.0'),
    (phi_uuid, null, 'dep/phi-ios', 3, '2.0'),
    (chi_uuid, null, 'dep/chi-ios', 3, '2.0'),
    (psi_uuid, null, 'dep/psi-ios', 3, '2.0'),
    (omega_uuid, null, 'dep/omega-ios', 3, '2.0')
  ;

  insert into dependencies
    (project_uuid, dependent_project_uuid, source_cd)
  values
    (zeta_uuid, phi_uuid, 0),
    (zeta_uuid, chi_uuid, 0),
    (zeta_uuid, psi_uuid, 0),
    (zeta_uuid, omega_uuid, 0),
    (phi_uuid, chi_uuid, 0),
    (phi_uuid, psi_uuid, 0)
  ;

end $$;
