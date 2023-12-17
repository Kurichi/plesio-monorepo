import 'package:kiikuten/domain/entity/tree_ranking.dart';
import 'package:kiikuten/domain/repository/tree_repository.dart';

class GetTreeRankingUseCase {
  final TreeRepository _treeRepository;

  GetTreeRankingUseCase(this._treeRepository);

  Future<List<TreeRanking>> execute() async {
    return await _treeRepository.getTreeRanking();
  }
}
