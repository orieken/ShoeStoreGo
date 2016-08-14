Feature: Cukes


  Scenario: Cukes are doubled
    Given i have 3 cukes
    When I use the doubler
    Then I should have 6 cukes
