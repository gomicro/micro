package features

const statusFeatureStub = `Feature:
  As a sys admin
  I want the service to have a status endpoint
  So I can tell if the service is working or not

  Scenario:
    Given a service
    When I hit the status endpoint
    Then I get a success response`
