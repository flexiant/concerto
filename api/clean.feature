Feature: No feature, just clean data if something went wrong

Scenario: Clean test data.
  When "automated-test1.concerto.io" domain is deleted if exists
  Then "automated-test1.concerto.io" domain doesn't exists
  When "automated-test2.concerto.io" domain is deleted if exists
  Then "automated-test2.concerto.io" domain doesn't exists
  When "automated-test3.concerto.io" domain is deleted if exists
  Then "automated-test3.concerto.io" domain doesn't exists
  When "automated-test4.concerto.io" domain is deleted if exists
  Then "automated-test4.concerto.io" domain doesn't exists
  When "automated-test5.concerto.io" domain is deleted if exists
  Then "automated-test5.concerto.io" domain doesn't exists
  When "automated-test6.concerto.io" domain is deleted if exists
  Then "automated-test6.concerto.io" domain doesn't exists
