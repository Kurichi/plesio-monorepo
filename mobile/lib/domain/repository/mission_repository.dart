import 'package:kiikuten/domain/entity/item.dart';
import 'package:kiikuten/domain/entity/mission.dart';

abstract class MissionRepository {
  Future<List<Mission>> getMissions();
  Future<UserMission> getUserMission(String userId);
  Future<void> updateUserMission(UserMission userMission);
}

class MissionRepositoryImpl implements MissionRepository {
  @override
  Future<List<Mission>> getMissions() async {
    return [
      Mission(
        id: '1',
        description: 'GitHubで1日に10回コミットする',
        term: '1日',
        target: 'GitHub',
        amount: 10,
        unit: '回',
        rewards: [
          Item(
            id: '1',
            name: 'ミツバチの群れ',
            description: 'ミツバチの群れを呼び寄せる',
            growthEffect: 10,
          ),
        ],
      ),
    ];
  }

  @override
  Future<UserMission> getUserMission(String userId) async {
    return UserMission(
      userId: userId,
      mission: Mission(
        id: '1',
        description: 'GitHubで1日に10回コミットする',
        term: '1日',
        target: 'GitHub',
        amount: 10,
        unit: '回',
        rewards: [
          Item(
            id: '1',
            name: 'ミツバチの群れ',
            description: 'ミツバチの群れを呼び寄せる',
            growthEffect: 10,
          ),
        ],
      ),
      progress: 0,
      deadline: 0,
    );
  }

  @override
  Future<void> updateUserMission(UserMission userMission) async {
    //
  }
}
