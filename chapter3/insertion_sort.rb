module LinkedList
  class InsertionSort
    def self.sort(list, &block)
      block = lambda { |val1, val2| val1 < val2 } unless block_given?
      return if list.empty?

      new_list = list.class.new
      iterator = list.start_node

      while iterator
        value_to_insert = iterator.value
        iterator = iterator.next_node
        node = new_list.start_node

        if node.nil? || block.call(value_to_insert, node.value)
          new_list.add_to_start(value_to_insert) && next
        end

        while node.next_node && !block.call(value_to_insert, node.next_node.value)
          node = node.next_node
        end

        new_list.add_after(node, value_to_insert)
      end

      new_list
    end
  end
end

# Example
# load 'linked_list.rb'
# list = LinkedList::UniDirectional.new
# list.add_to_start(5)
# list.add_to_start(9)
# list.add_to_start(3)
# list.add_to_start(1)
# list.add_to_start(2)
# LinkedList::InsertionSort.sort(list).each { |v| p v.value }
# LinkedList::InsertionSort.sort(list) { |node, node2| node > node2 }.each { |v| p v.value }
