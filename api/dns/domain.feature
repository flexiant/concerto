Feature: User manage DNS using concerto
  In order to offer the user a handy DNS tool
  As a concerto user with permissions granted to manage DNS
  I want be able to create, list, update and delete domains and records

  Scenario: create a simple domain
    Given "automated-test1.concerto.io" domain doesn't exists
    When "automated-test1.concerto.io" domain is created with contact "test@automated-test1.concerto.io"
    Then "automated-test1.concerto.io" domain exists
    When "automated-test1.concerto.io" domain is deleted
    Then "automated-test1.concerto.io" domain doesn't exists

  Scenario: update a simple domain
    Given "automated-test2.concerto.io" domain doesn't exists
    When "automated-test2.concerto.io" domain is created with contact "test@automated-test2.concerto.io"
    Then "automated-test2.concerto.io" domain exists
    When "automated-test2.concerto.io" domain is updated with contact "updated@automated-test2.concerto.io"
    Then list domains should include:
      | name                        | contact                           |
      | automated-test2.concerto.io | updated@automated-test2.concerto.io  |
    When "automated-test2.concerto.io" domain is deleted
    Then "automated-test2.concerto.io" domain doesn't exists

  Scenario: list domains
    Given "automated-test3.concerto.io" domain doesn't exists
    And "automated-test4.concerto.io" domain doesn't exists
    And "automated-test5.concerto.io" domain doesn't exists
    When "automated-test3.concerto.io" domain is created with contact "test@automated-test3.concerto.io"
    And "automated-test4.concerto.io" domain is created with contact "test@automated-test4.concerto.io"
    And "automated-test5.concerto.io" domain is created with contact "test@automated-test5.concerto.io"
    Then list domains should include:
      | name                        | contact                           |
      | automated-test3.concerto.io | test@automated-test3.concerto.io  |
      | automated-test4.concerto.io | test@automated-test4.concerto.io  |
      | automated-test5.concerto.io | test@automated-test5.concerto.io  |
    When "automated-test3.concerto.io" domain is deleted
    Then "automated-test3.concerto.io" domain doesn't exists
    When "automated-test4.concerto.io" domain is deleted
    Then "automated-test4.concerto.io" domain doesn't exists
    When "automated-test5.concerto.io" domain is deleted
    Then "automated-test5.concerto.io" domain doesn't exists

  Scenario: create domain record
    Given "automated-test6.concerto.io" domain doesn't exists
    When "automated-test6.concerto.io" domain is created with contact "test@automated-test7.concerto.io"
    Then "automated-test6.concerto.io" domain exists
    When records for domain "automated-test6.concerto.io" are created:
    | type  | name                | content   | server_id | prio |
    | A     | automated-test6-1   | 10.0.0.1  |           |      |
    Then records for domain "automated-test6.concerto.io" contain:
    | type  | name                | content   | server_id | prio |
    | A     | automated-test6-1   | 10.0.0.1  |           |      |    
    When "automated-test6.concerto.io" domain is deleted
    Then "automated-test6.concerto.io" domain doesn't exists
