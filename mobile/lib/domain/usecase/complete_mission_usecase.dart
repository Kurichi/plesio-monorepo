import 'package:kiikuten/domain/repository/mission_repository.dart';

class CompleteMissionUseCase {
  final MissionRepository missionRepository;

  CompleteMissionUseCase(this.missionRepository);

  Future<void> execute(String userId, String missionId) async {
    var userMission = await missionRepository.getUserMission(userId);

    if (userMission.mission.id == missionId) {
      userMission.progress = userMission.mission.amount;
      await missionRepository.updateUserMission(userMission);
    }
  }
}
