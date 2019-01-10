module LinkedList
  class UniDirectional
    class Node
      attr_reader :value
      attr_accessor :next_node

      def initialize(value)
        @value = value
      end

      def next_node=(node)
        raise ArgumentError if !node.nil? && !node.is_a?(self.class)

        @next_node = node
      end
    end

    attr_accessor :start_node

    def each
      iterator = start_node
      while iterator
        yield iterator
        iterator = iterator.next_node
      end
    end

    def add_to_start(value)
      node = ensure_node(value)
      node.next_node = start_node unless list_empty?
      @start_node = node
    end

    def add_to_end(value)
      add_to_start && return if list_empty?

      current_node = start_node

      while current_node.next_node
        current_node = current_node.next_node
      end

      current_node.next_node = ensure_node(value)
    end

    def add_after(node, value)
      return false unless node.is_a?(Node)
      return false unless node_in_list?(node)

      new_node  = ensure_node(value)

      new_node.next_node = node.next_node
      node.next_node     = new_node

      new_node
    end

    def delete_after(node)
      return false unless node.is_a?(Node)
      return false unless node_in_list?(node)
      return false if last_node?(node)

      node.next_node = node.next_node.next_node

      true
    end

    def delete_node(node)
      return false unless node.is_a?(Node)
      return false unless node_in_list?(node)

      if first_node?(node)
        @start_node = node.next_node
        return true
      end

      iterator = start_node
      while iterator.next_node
        if iterator.next_node == node
          iterator.next_node = node.next_node
          return true
        end

        iterator = iterator.next_node
      end

      false
    end

    def empty?
      start_node.nil?
    end

    private

      def node_in_list?(node)
        iterator = start_node
        while iterator
          return true if iterator == node
          iterator = iterator.next_node
        end

        false
      end

      def last_node?(node)
        node.next_node.nil?
      end

      def first_node?(node)
        node == start_node
      end

      def ensure_node(value)
        return value if value.is_a?(Node)

        Node.new(value)
      end

      def list_empty?
        @start_node.nil?
      end
  end
end
