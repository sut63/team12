/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntClubEdges,
    EntClubEdgesFromJSON,
    EntClubEdgesFromJSONTyped,
    EntClubEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntClub
 */
export interface EntClub {
    /**
     * 
     * @type {number}
     * @memberof EntClub
     */
    clubBranchID?: number;
    /**
     * 
     * @type {number}
     * @memberof EntClub
     */
    clubTypeID?: number;
    /**
     * 
     * @type {EntClubEdges}
     * @memberof EntClub
     */
    edges?: EntClubEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntClub
     */
    id?: number;
    /**
     * Name holds the value of the "name" field.
     * @type {string}
     * @memberof EntClub
     */
    name?: string;
    /**
     * Purpose holds the value of the "purpose" field.
     * @type {string}
     * @memberof EntClub
     */
    purpose?: string;
    /**
     * 
     * @type {number}
     * @memberof EntClub
     */
    userID?: number;
}

export function EntClubFromJSON(json: any): EntClub {
    return EntClubFromJSONTyped(json, false);
}

export function EntClubFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntClub {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'clubBranchID': !exists(json, 'clubBranch_ID') ? undefined : json['clubBranch_ID'],
        'clubTypeID': !exists(json, 'clubType_ID') ? undefined : json['clubType_ID'],
        'edges': !exists(json, 'edges') ? undefined : EntClubEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'purpose': !exists(json, 'purpose') ? undefined : json['purpose'],
        'userID': !exists(json, 'userID') ? undefined : json['userID'],
    };
}

export function EntClubToJSON(value?: EntClub | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'clubBranch_ID': value.clubBranchID,
        'clubType_ID': value.clubTypeID,
        'edges': EntClubEdgesToJSON(value.edges),
        'id': value.id,
        'name': value.name,
        'purpose': value.purpose,
        'userID': value.userID,
    };
}


