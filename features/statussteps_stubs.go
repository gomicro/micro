package features

const statusStepsStub = `# Givens
Given('a service') do
  # noop, running the service is handled outside of this
end

# Whens
When('I hit the status endpoint') do
  @response = get('/v1/status')
end

# Thens
Then('I get a success response') do
  expect(@response).not_to(be_nil, 'expected: response not nil\ngot: response nil')
  expect(@response.code.to_i).to(eql(200))
end`
