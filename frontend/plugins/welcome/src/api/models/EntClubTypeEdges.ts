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
    EntClub,
    EntClubFromJSON,
    EntClubFromJSONTyped,
    EntClubToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntClubTypeEdges
 */
export interface EntClubTypeEdges {
    /**
     * Club holds the value of the club edge.
     * @type {Array<EntClub>}
     * @memberof EntClubTypeEdges
     */
    club?: Array<EntClub>;
}

export function EntClubTypeEdgesFromJSON(json: any): EntClubTypeEdges {
    return EntClubTypeEdgesFromJSONTyped(json, false);
}

export function EntClubTypeEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntClubTypeEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'club': !exists(json, 'club') ? undefined : ((json['club'] as Array<any>).map(EntClubFromJSON)),
    };
}

export function EntClubTypeEdgesToJSON(value?: EntClubTypeEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'club': value.club === undefined ? undefined : ((value.club as Array<any>).map(EntClubToJSON)),
    };
}


