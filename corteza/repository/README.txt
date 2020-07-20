Goal
 - multi-repository (different DB) support
 - refactor old raw-sql migrations to logical and dynamic schema provisioning
 - (tbd) consolidation of system & compose (& messaging) subsystems

File system:
 /repository            Holds all repository logic for all subsystems, for all core repository implementation
   /internal            Internal repository tools (pkg/ql, pgk/rh should be moved here)
 /<implementation>      Individual core repository implementation
                        [mysql|postgresql|redis|memory|sqlite|elasticsearch|mongo]
   /schema              Schema provisioning for individual repository implementation
