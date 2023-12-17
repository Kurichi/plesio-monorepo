import 'package:kiikuten/domain/repository/tree_repository.dart';
import 'package:kiikuten/provider/repository/tree_repository_provider.dart';
import 'package:riverpod_annotation/riverpod_annotation.dart';

part 'plant_tree_usecase.g.dart';

@riverpod
PlantTreeUseCase plantTreeUseCase(PlantTreeUseCaseRef ref) {
  final treeRepository = ref.watch(treeRepositoryProvider);
  return PlantTreeUseCase(treeRepository);
}

class PlantTreeUseCase {
  final TreeRepository _treeRepository;

  PlantTreeUseCase(this._treeRepository);

  Future<void> execute() async {
    await _treeRepository.plantTree();
  }
}
