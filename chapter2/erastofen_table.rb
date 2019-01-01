class ErastofenTable
  FIRST_PRIME_NUMBER = 2
  attr_reader :number, :range_mask, :upper_limit

  def initialize(number)
    @number      = number
    @range_mask  = Array.new(number, false)
    @upper_limit = Math.sqrt(number)
  end

  def call
    return if number < FIRST_PRIME_NUMBER

    compose_mask
    apply_mask
  end

  private

  def compose_mask
    current_number = FIRST_PRIME_NUMBER

    while current_number <= upper_limit
      (current_number * 2 .. number).step(current_number).each do |number|
        range_mask[number - 1] = true
      end
      current_number = get_next_number(current_number)
    end
  end

  def apply_mask
    range_mask.map.with_index do |is_composite, index|
      next if is_composite || index == 0

      index + 1
    end.compact
  end

  def get_next_number(current_number)
    range_mask.each_with_index do |is_composite, index|
      next if is_composite || index + 1 <= current_number

      return index + 1
    end
  end
end
