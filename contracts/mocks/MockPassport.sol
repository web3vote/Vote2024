//SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "../Passport.sol";

contract MockPassport is Passport {
    // Open internal funcs for test
    function createUser(
        address user_address,
        PassportType id_type,
        bytes32 mrz_uid_hash
    ) public {
        super._createUser(user_address, id_type, mrz_uid_hash);
    }

    function create_and_proove_ttp(
        address user_address,
        PassportType id_type,
        bytes32 mrz_uid_hash
    ) public {
        super._create_and_proove_ttp(user_address, id_type, mrz_uid_hash);
    }

    function proove_ttp(bytes32 mrz_uid_hash) public {
        super._proove_ttp(mrz_uid_hash);
    }

    function getTTP_checks_by_service(address service) public view {
        super._getTTP_checks_by_service(service);
    }

    function getTTP_proofs_of_user(bytes32 mrz_uid_hash) public view {
        super._getTTP_proofs_of_user(mrz_uid_hash);
    }
}
