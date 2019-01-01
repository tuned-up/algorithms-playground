require 'forwardable'

class ErastofenTable
  class IntegerWithInfo
    extend Forwardable

    attr_reader :number
    attr_accessor :is_composite

    def_delegators :@number, :<=>, :+, :-, :*

    def initialize(number)
      @number       = number
      @is_composite = false
    end
    alias is_composite? is_composite

    def mark_as_composite!
      @is_composite = true
    end

    def to_int
      @number
    end
  end

  FIRST_PRIME_NUMBER = 2
  attr_reader :number, :range, :upper_limit

  def initialize(number)
    @number      = number
    @range       = (1..number).map { |number| IntegerWithInfo.new(number) }
    @upper_limit = Math.sqrt(number)
  end

  def call
    return if number < FIRST_PRIME_NUMBER

    current_number = FIRST_PRIME_NUMBER

    while current_number <= upper_limit
      (current_number * 2 .. number).step(current_number).each do |number|
        range[number - 1].mark_as_composite!
      end
      current_number = get_next_number(current_number)
    end

    range[1..-1].reject { |number| number.is_composite || number == 1 }.map(&:to_int)
  end

  private

  def get_next_number(current_number)
    range.each do |number|
      next if number.is_composite? || number <= current_number.to_int

      return number
    end
  end
end
