import 'package:kiikuten/domain/repository/mission_repository.dart';

class ProgressMissionUseCase {
  final MissionRepository _missionRepository;

  ProgressMissionUseCase(this._missionRepository);

  Future<void> execute(String missionId) async {
    await _missionRepository.progressMission(missionId);
  }
}
