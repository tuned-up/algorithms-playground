class FermaPrimeTest
  attr_reader :number, :num_tests

  def self.is_prime?(number, num_tests = 10)
    new(number, num_tests).call
  end

  def initialize(number, num_tests = 10)
    @number    = number
    @num_tests = num_tests
  end

  def call
    num_tests.times do
      n = rand(1...number)
      return false unless (n ** (number - 1) % number == 1)
    end

    true
  end
end
