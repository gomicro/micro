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
Then('I get a sucess response') do
	expect(@response.code.to_i).to(eql(200))
end`
