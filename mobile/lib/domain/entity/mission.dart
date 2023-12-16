import 'package:kiikuten/domain/entity/item.dart';

class Mission {
  final String id;
  final String description;
  final String term;
  final String target;
  final int amount;
  final String unit;
  final List<Item> rewards;

  Mission({
    required this.id,
    required this.description,
    required this.term,
    required this.target,
    required this.amount,
    required this.unit,
    required this.rewards,
  });
}

class UserMission {
  final String userId;
  final Mission mission;
  int progress;
  final int deadline;

  UserMission({
    required this.userId,
    required this.mission,
    required this.progress,
    required this.deadline,
  });

  bool isCompleted() {
    return progress >= mission.amount;
  }
}
