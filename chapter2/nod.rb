class Nod
  attr_reader :first_number, :second_number, :result

  def initialize(first_number, second_number)
    @first_number, @second_number = [first_number, second_number].map(&:to_int).sort
  end

  def call
    return false if
    while second_number > 0
      remainder = first_number % second_number

      @first_number = second_number
      @second_number = remainder
    end

    @result = first_number
  end
end
