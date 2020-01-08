package benchs

import (
	"github.com/dshulyak/urkeltrie"
	"github.com/dshulyak/urkeltrie/store"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb/leveldb"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/tendermint/iavl"
	db "github.com/tendermint/tm-db"
)

type testTree interface {
	Write([]byte, []byte) error
	Commit() error
}

func newPatricia(path string) (testTree, error) {
	ldb, err := leveldb.New(path, 32, 32, "metrics")
	if err != nil {
		return nil, err
	}

	db := trie.NewDatabaseWithCache(ldb, 32)
	trie, err := trie.New(common.Hash{}, db)
	if err != nil {
		return nil, err
	}
	return &patricia{
		db:     ldb,
		trie:   trie,
		trieDB: db,
	}, nil
}

type patricia struct {
	db     *leveldb.Database
	trie   *trie.Trie
	trieDB *trie.Database
}

func (p *patricia) Write(key, value []byte) error {
	p.trie.Update(key, value)
	return nil
}

func (p *patricia) Commit() error {
	hash, err := p.trie.Commit(nil) // accepts leaf callback
	if err != nil {
		return err
	}
	return p.trieDB.Commit(hash, false)
}

func newIavlTree(path string) (testTree, error) {
	db := db.NewDB("bench", db.GoLevelDBBackend, path)
	tree := iavl.NewMutableTree(db, 10000)
	return &iavlTree{tree: tree}, nil
}

type iavlTree struct {
	tree *iavl.MutableTree
}

func (it *iavlTree) Write(key, value []byte) error {
	it.tree.Set(key, value)
	return nil
}

func (it *iavlTree) Commit() error {
	_, _, err := it.tree.SaveVersion()
	return err
}

func newUrkel(path string) (testTree, error) {
	conf := store.DefaultProdConfig(path)
	conf.TreeWriteBuffer = 64 << 20
	fs, err := store.Open(conf)
	if err != nil {
		return nil, err
	}
	return &urkel{tree: urkeltrie.NewTree(fs)}, nil
}

type urkel struct {
	tree *urkeltrie.Tree
}

func (u *urkel) Write(key, value []byte) error {
	return u.tree.Put(key, value)
}

func (u *urkel) Commit() error {
	return u.tree.Commit()
}
