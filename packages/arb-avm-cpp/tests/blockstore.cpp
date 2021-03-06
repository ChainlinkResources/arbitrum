//
//  blockstore.cpp
//  avm
//
//  Created by Harry Kalodner on 5/17/20.
//

#include "config.hpp"
#include "helper.hpp"

#include <data_storage/blockstore.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/storageresult.hpp>

#include <catch2/catch.hpp>

TEST_CASE("BlockStore tests") {
    DBDeleter deleter;
    DataStorage storage{dbpath};
    auto store = storage.getBlockStore();

    SECTION("BlockStore min, max, and empty") {
        REQUIRE(store->isEmpty());
        REQUIRE(store->maxHeight() == 0);
        REQUIRE(store->minHeight() == 0);

        REQUIRE(store->putBlock(10, 30, {1, 2, 3}).ok());
        REQUIRE(!store->isEmpty());
        REQUIRE(store->maxHeight() == 10);
        REQUIRE(store->minHeight() == 10);

        REQUIRE(store->putBlock(20, 30, {1, 2, 3}).ok());
        REQUIRE(!store->isEmpty());
        REQUIRE(store->maxHeight() == 20);
        REQUIRE(store->minHeight() == 10);

        REQUIRE(store->putBlock(5, 30, {1, 2, 3}).ok());
        REQUIRE(!store->isEmpty());
        REQUIRE(store->maxHeight() == 20);
        REQUIRE(store->minHeight() == 5);

        REQUIRE(store->putBlock(15, 30, {1, 2, 3}).ok());
        REQUIRE(!store->isEmpty());
        REQUIRE(store->maxHeight() == 20);
        REQUIRE(store->minHeight() == 5);

        REQUIRE(store->deleteBlock(20, 30).ok());
        REQUIRE(store->maxHeight() == 15);
        REQUIRE(store->minHeight() == 5);

        REQUIRE(store->deleteBlock(5, 30).ok());
        REQUIRE(store->maxHeight() == 15);
        REQUIRE(store->minHeight() == 10);
    }

    SECTION("BlockStore block hashes at height") {
        REQUIRE(store->blockHashesAtHeight(10).empty());

        REQUIRE(store->putBlock(10, 30, {1, 2, 3}).ok());
        REQUIRE(store->blockHashesAtHeight(10) == std::vector<uint256_t>{30});
        REQUIRE(store->blockHashesAtHeight(9).empty());
        REQUIRE(store->blockHashesAtHeight(11).empty());

        REQUIRE(store->putBlock(10, 40, {1, 2, 3}).ok());
        REQUIRE(store->blockHashesAtHeight(10) ==
                std::vector<uint256_t>{30, 40});

        REQUIRE(store->deleteBlock(10, 30).ok());
        REQUIRE(store->blockHashesAtHeight(10) == std::vector<uint256_t>{40});
    }

    SECTION("BlockStore put and get") {
        auto result = store->getBlock(10, 30);
        REQUIRE(!result.status.ok());
        REQUIRE(result.data.empty());

        REQUIRE(store->putBlock(10, 30, {1, 2, 3}).ok());
        result = store->getBlock(10, 30);
        REQUIRE(result.status.ok());
        REQUIRE(result.data == std::vector<unsigned char>{1, 2, 3});

        REQUIRE(store->putBlock(10, 30, {2, 2, 3, 4}).ok());
        result = store->getBlock(10, 30);
        REQUIRE(result.status.ok());
        REQUIRE(result.data == std::vector<unsigned char>{2, 2, 3, 4});

        REQUIRE(store->deleteBlock(10, 30).ok());
        result = store->getBlock(10, 30);
        REQUIRE(!result.status.ok());
        REQUIRE(result.data.empty());
    }
}
