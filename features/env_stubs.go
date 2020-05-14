package features

const envStub = `# frozen_string_literal: true
require 'json'
require 'net/http'

module Cucumber
  @base_url = 'http://localhost/v1'

  def self.base_url(url = nil)
    @base_url = url.sub(%r{\/+$}, '') unless url.nil?
    return @base_url
  end

  def self.default_env(key, default)
    return ENV.key?(key) ? ENV[key] : default
  end

  def self.bearer_token(token = nil)
    @bearer_token = token unless token.nil?
    auth_scheme('bearer')
    return @bearer_token
  end

  def self.basic_auth(username = nil, password = nil)
    @username = username unless username.nil?
    @password = password unless password.nil?
    auth_scheme('basic')
    return @username, @password
  end

  def self.auth_scheme(scheme = nil)
    scheme = scheme.downcase unless scheme.nil?
    raise ArgumentError, 'auth scheme must be one of basic or bearer' unless ['basic', 'bearer', nil].include?(scheme)
    @auth_scheme = scheme unless scheme.nil?
    return @auth_scheme
  end

  def self.headers=(headers)
    @headers = headers
  end

  def self.headers
    @headers || {}
  end

  def self.do_req(endpoint, method, body)
    endpoint = endpoint.sub(%r{^\/+}, '')
    uri = URI("#{@base_url}/#{endpoint}")
    http = Net::HTTP.new(uri.host, uri.port)
    http.use_ssl = uri.scheme == 'https'
    http.verify_mode = OpenSSL::SSL::VERIFY_NONE
    http.ca_file = '/etc/ssl/certs/ca-certificates.crt'

    req = method.new(uri)
    headers.each do |name, value|
      req.add_field(name, value)
    end
    req.add_field('Authorization', "Bearer #{@bearer_token}") if auth_scheme == 'bearer'
    req.basic_auth(@username, @password) if auth_scheme == 'basic'
    req.add_field('Content-Type', 'application/json')
    req.body = body.to_json unless body.nil?
    resp = http.request(req)

    return resp
  end

  def self.delete(endpoint, body)
    return do_req(endpoint, Net::HTTP::Delete, body)
  end

  def self.get(endpoint)
    return do_req(endpoint, Net::HTTP::Get, nil)
  end

  def self.patch(endpoint, body)
    return do_req(endpoint, Net::HTTP::Patch, body)
  end

  def self.post(endpoint, body)
    return do_req(endpoint, Net::HTTP::Post, body)
  end

  def self.put(endpoint, body)
    return do_req(endpoint, Net::HTTP::Put, body)
  end
end

BASEURL = Cucumber.default_env('{{ .Name }}_CUCUMBER_BASEURL', 'http://app:4567/')
Cucumber.base_url(BASEURL)`
