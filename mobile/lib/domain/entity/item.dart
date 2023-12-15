class Item {
  final String id;
  final String name;
  final int growthEffect;
  bool isUsed = false;

  Item({
    required this.id,
    required this.name,
    required this.growthEffect,
  });

  void use() {
    isUsed = true;
  }
}
