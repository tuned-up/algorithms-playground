class Disk
  attr_reader :d
  def initialize(d)
    @d = d
  end

  def smaller?(other_disk)
    self.d < other_disk.d
  end
end

class Stick
  attr_reader :disks
  def initialize(disk_count = 0)
    @disks = []
    disk_count.downto(1) do |n|
      place(Disk.new(n))
    end
  end

  def lift
    disks.pop
  end

  def place(disk)
    if disks.last&.smaller?(disk)
      raise "cannot place"
    end

    disks << disk
  end
end

class HanoiTower
  attr_reader :sticks
  def initialize(sticks)
    @sticks = sticks
  end

  def move(from, to, buf, num, &block)
    if num > 1
      move(from, buf, to, num -1, &block)
    end

    to.place(from.lift)
    block.call if block

    if num > 1
      move(buf, to, from, num - 1, &block)
    end
  end

  def ugly_print
    sticks.each do |s|
      print('|')
      s.disks.each do |d|
        print(d.d)
      end
      print('-')
      puts
    end

    puts
  end
end

class HanoiTowerSolver
  attr_reader :tower
  PAUSE_TIME = 0.4

  def initialize(tower)
    @tower = tower
  end

  def solve
    tower.move(tower.sticks[0], tower.sticks[1], tower.sticks[2], tower.sticks[0].disks.size) do
      tower.ugly_print
      sleep(PAUSE_TIME)
    end
  end
end

tower = HanoiTower.new([Stick.new(5), Stick.new, Stick.new])
tower.ugly_print
HanoiTowerSolver.new(tower).solve
tower.ugly_print
