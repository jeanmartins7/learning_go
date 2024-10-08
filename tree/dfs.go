package tree

func dfs_search_recursive(node *Node, target int) bool {

	if node == nil {
		return false
	}

	if node.Value == target {
		return true
	}

	if dfs_search_recursive(node.Left, target) {
		return true
	}

	if dfs_search_recursive(node.Right, target) {
		return true
	}

	return false
}
