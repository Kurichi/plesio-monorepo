class ItemModel {
  final String id;
  final String name;
  final String description;
  final int growthEffect;

  ItemModel({
    required this.id,
    required this.name,
    required this.description,
    required this.growthEffect,
  });

  factory ItemModel.fromJson(Map<String, dynamic> json) {
    return ItemModel(
      id: json['id'],
      name: json['name'],
      description: json['description'],
      growthEffect: json['growthEffect'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'description': description,
      'growthEffect': growthEffect,
    };
  }
}
