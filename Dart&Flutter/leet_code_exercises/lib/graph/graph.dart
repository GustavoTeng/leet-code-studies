class Graph{
  int numBusesToDestination(List<List<int>> routes, int source, int target) {
    if (source == target) {
      return 0;
    }

    Map<int, List<int>> graph = {};
    for (int route = 0; route < routes.length; route++) {
      for (int stop = 0; stop < routes[route].length; stop++) {
        int busStop = routes[route][stop];
        int stopLastItem = stop + 1 == routes[route].length ? 0 : stop + 1;
        if (graph[busStop] == null) {
          graph[busStop] = [routes[route][stopLastItem]];
        } else {
          graph[busStop]!.add(routes[route][stopLastItem]);
        }
      }
    }
    print(graph);
    
    Map<int,int> queue = {};
    List<int> backTrack = [];
    
    if (graph.containsKey(source)) {
          for(var item in graph[source]!){
            queue.putIfAbsent(item ,() => 1);
            backTrack.add(item);
          }
      }
    
    while(queue.length > 0){
      var firstKey = queue.keys.first;
      var firstItem = queue.remove(firstKey);
      int busCounter = 0;
      print(firstItem);
    }

    return -1;
  }

  alienDictionary(List<String> words){
    Map<String, Set<String>> charNodes = {};
    for (String word in words) {
      for (String char in word.split('')) {
        charNodes.putIfAbsent(char, () => {});
      }
    }
    
    for (int i = 0; i < words.length - 1; i++) {
      String w1 = words[i];
      String w2 = words[i + 1];
      int minLen = w1.length < w2.length ? w1.length : w2.length;
      
      if (w1.length > w2.length && w1.substring(0, minLen) == w2.substring(0, minLen)) {
        return '';
      }

      for (int j = 0; j < minLen; j++) {
        String char1 = w1[j];
        String char2 = w2[j];

        if (char1 != char2) {
          charNodes[char1]!.add(char2);
          break;
        }
      }
    }
    
    Map<String, bool> visited = {};
    List<String> result = [];

    bool dfs(String char) {
      if (visited[char] == true) return true;
      if (visited[char] == false) return false;

      visited[char] = true;

      for (final String neighChar in charNodes[char]!) {
        if (dfs(neighChar)) {
          return true;
        }
      }

      visited[char] = false;
      result.add(char);
      return false;
    }

    for (final String char in charNodes.keys) {
      if (dfs(char)) {
        return '';
      }
    }

    return result.reversed.join('');
  }
}