import 'package:kiikuten/domain/repository/mission_repository.dart';
import 'package:kiikuten/provider/repository/mission_repository_provider.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'progress_mission_usecase.g.dart';

@riverpod
ProgressMissionUseCase progressMissionUseCase(ProgressMissionUseCaseRef ref) {
  final missionRepository = ref.watch(missionRepositoryProvider);
  return ProgressMissionUseCase(missionRepository);
}

class ProgressMissionUseCase {
  final MissionRepository _missionRepository;

  ProgressMissionUseCase(this._missionRepository);

  Future<void> execute(String missionId) async {
    await _missionRepository.progressMission(missionId);
  }
}
