require_relative 'linked_list'

module CycleBreaker
  class WithHash
    attr_reader :linked_list

    def initialize(linked_list)
      @linked_list = linked_list
    end

    def break_cycle!
      node = linked_list.start_node
      nodes_seen = {}

      while node.next_node && !nodes_seen[node.next_node]
        nodes_seen[node] = true
        node = node.next_node
      end
      node.next_node = nil

      linked_list
    end
  end
end

# Example
# load 'cycle_breaker.rb'
# list = LinkedList::UniDirectional.new
# list.add_to_start(5)
# list.add_to_end(9)
# node = list.add_to_end(3)
# list.add_to_end(1)
# node_2 = list.add_to_end(2)
# node_2.next_node = node
# CycleBreaker::WithHash.new(list).break_cycle!
