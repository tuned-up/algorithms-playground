class PrimeDividers
  attr_accessor :number, :factors

  def initialize(number)
    @number  = number
    @factors = []
  end

  def call
    max_factor = Math.sqrt(number)
    current_factor = 2

    while current_factor <= max_factor
      while number % current_factor == 0
        factors << current_factor
        @number /= current_factor
        max_factor = Math.sqrt(number)
      end

      current_factor += 1
    end

    factors << number if number > 1

    factors
  end
end
