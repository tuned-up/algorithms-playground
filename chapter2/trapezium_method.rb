class TrapeziumMethod
  attr_reader :xmin, :xmax, :num_intervals, :block

  def initialize(xmin, xmax, num_intervals, &block)
    @xmin = xmin
    @xmax = xmax
    @num_intervals = num_intervals
    @block = block
  end

  def call
    delta_x    = (xmax - xmin) / num_intervals.to_f
    total_area = 0
    current_x  = xmin

    num_intervals.times do |interval_num|
      total_area += delta_x * ((block.call(current_x) + block.call(current_x + delta_x)) / 2)
      current_x  += delta_x
    end

    total_area
  end
end
